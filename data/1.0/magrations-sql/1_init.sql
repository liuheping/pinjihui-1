create type address as
(
  province_id integer,
  city_id     integer,
  area_id     integer,
  region_name varchar(255),
  address     varchar(255)
);

create type favorite_type as enum ('product', 'store');

create type how_oos as enum ('together', 'cancel', 'consult');

create type input_types as enum ('textfield', 'textarea', 'dropdown', 'mutiselect');

create type order_status as enum ('unconfirmed', 'confirmed', 'cancelled', 'invalid', 'returned');

create type pay_status as enum ('unpaid', 'paying', 'paid');

create type roles as enum ('admin', 'customer', 'provider', 'ally');

create type shipping_status as enum ('unshipped', 'shipped', 'invalid', 'returned');

create type shipping_address as
(
  consignee varchar(255),
  address   address,
  zipcode   varchar(32),
  mobile    varchar(32)
);

create type debit_card as
(
  card_holder varchar(255),
  bank_name   varchar(255),
  card_number varchar(32),
  province    integer,
  city        integer,
  branch      varchar(255)
);

create type cash_request_status as enum ('unchecked', 'checking', 'checked', 'paid', 'finished', 'refused', 'closed');

create table orders
(
  id              varchar(32)                         not null
    constraint orders_pkey
    primary key,
  user_id         varchar(32)                         not null,
  order_status    order_status                        not null,
  shipping_status shipping_status                     not null,
  pay_status      pay_status                          not null,
  postscript      varchar(255),
  shipping_id     varchar(32)                         not null,
  shipping_name   varchar(255)                        not null,
  pay_id          varchar(32)                         not null,
  pay_name        varchar(255)                        not null,
  how_oos         how_oos,
  inv_payee       varchar(32),
  inv_type        varchar(255),
  inv_content     varchar(255),
  amount          money                               not null,
  shipping_fee    money,
  pay_fee         money,
  money_paid      money,
  order_amount    money,
  created_at      timestamp default CURRENT_TIMESTAMP not null,
  confirm_time    timestamp,
  pay_time        timestamp,
  shipping_time   timestamp,
  tax             money,
  parent_id       varchar(32),
  merchant_id     varchar(32),
  address         shipping_address                    not null,
  note            text
);

comment on table orders
is '订单表';

comment on column orders.id
is '订单号';

comment on column orders.order_status
is '订单状态';

comment on column orders.shipping_status
is '商品配送情况';

comment on column orders.pay_status
is '支付状态';

comment on column orders.postscript
is '订单留言';

comment on column orders.shipping_id
is '配送方式ID';

comment on column orders.how_oos
is '缺货/备货处理方式';

comment on column orders.inv_payee
is '发票抬头';

comment on column orders.inv_type
is '发票类型';

comment on column orders.inv_content
is '发票内容，取值配置表';

comment on column orders.amount
is '商品总价格';

comment on column orders.shipping_fee
is '配送费用';

comment on column orders.pay_fee
is '支付费用,跟支付方式的配置相关,取值表payment';

comment on column orders.money_paid
is '已付款金额';

comment on column orders.order_amount
is '应付款金额';

comment on column orders.created_at
is '订单生成时间';

comment on column orders.confirm_time
is '订单确认时间';

comment on column orders.pay_time
is '支付时间';

comment on column orders.shipping_time
is '订单配送时间';

comment on column orders.tax
is '税额';

comment on column orders.merchant_id
is '商家ID';

create table products
(
  id               varchar(32)                         not null
    constraint products_pkey
    primary key,
  name             varchar(255)                        not null,
  is_sale          boolean default false               not null,
  attribute_set_id varchar(32),
  batch_price      money                               not null,
  second_price     money                               not null,
  retail_price     money                               not null,
  category_id      varchar(32),
  related_ids      character varying [],
  content          text,
  brand_id         varchar(32),
  deleted          boolean default false               not null,
  created_at       timestamp default CURRENT_TIMESTAMP not null,
  updated_at       timestamp default CURRENT_TIMESTAMP not null,
  tags             character varying [],
  attrs            jsonb,
  recommended      boolean default false               not null
);

comment on table products
is '商品表';

comment on column products.batch_price
is '批发价,即供货商给平台的价格';

comment on column products.second_price
is '平台给加盟商的价格';

comment on column products.retail_price
is '零售价';

comment on column products.related_ids
is '关联商品ID';

comment on column products.content
is '商品详情';

comment on column products.brand_id
is '品牌';

comment on column products.tags
is '标签';

comment on column products.attrs
is '属性';

create table regoins
(
  id         serial              not null
    constraint regoins_pkey
    primary key,
  parent_id  integer default 0   not null,
  name       varchar(32)         not null,
  sort_order integer default 255 not null
);

comment on table regoins
is '行政区域表';

create table users
(
  id              varchar(32)                            not null
    constraint users_pkey
    primary key,
  name            varchar(32),
  mobile          varchar(32),
  password        bytea                                  not null,
  type            roles                                  not null,
  email           varchar(255),
  created_at      timestamp(6) default CURRENT_TIMESTAMP not null,
  updated_at      timestamp(6) default CURRENT_TIMESTAMP not null,
  status          smallint default 0                     not null,
  last_ip         inet,
  last_login_time timestamp(6)
);

comment on table users
is '用户表';

comment on column users.status
is '用户状态,0:正常, 1:禁用';

comment on column users.last_login_time
is '最近一次登录时间';

create table addresses
(
  id         varchar(32)           not null
    constraint addresses_pkey
    primary key,
  user_id    varchar(32)           not null
    constraint addresses_users_id_fk
    references users,
  consignee  varchar(32)           not null,
  address    address               not null,
  zipcode    varchar(6),
  mobie      varchar(11),
  is_default boolean default false not null
);

comment on table addresses
is '用户收货地址表';

comment on column addresses.consignee
is '收货人姓名';

create table rel_merchants_products
(
  product_id varchar(32)       not null
    constraint rel_merchants_products_products_id_fk
    references products,
  user_id    varchar(32)       not null
    constraint rel_merchants_products_users_id_fk
    references users,
  stock      integer default 1 not null,
  constraint rel_merchants_products_user_id_product_id_pk
  primary key (user_id, product_id)
);

comment on table rel_merchants_products
is '商家商品关联表';

comment on column rel_merchants_products.stock
is '库存';

create table wecharts
(
  openid      varchar(32)        not null
    constraint wecharts_pkey
    primary key,
  session_key varchar(32)        not null,
  nick_name   varchar(255)       not null,
  gender      smallint default 0 not null,
  language    varchar(32),
  city        varchar(32),
  province    varchar(32),
  country     varchar(32),
  avatar_url  varchar(526),
  user_id     varchar(32)
);

comment on table wecharts
is '微信账号信息表';

comment on column wecharts.gender
is '1:男,2:女,0未知';

create table categories
(
  id         varchar(32)  not null
    constraint categories_pkey
    primary key,
  parent_id  varchar(32)  not null,
  name       varchar(255) not null,
  sort_order integer default 255
);

comment on table categories
is '分类表';

comment on column categories.sort_order
is '排序号';

create table attributes
(
  id            varchar(255)          not null
    constraint attributes_pkey
    primary key,
  name          varchar(255)          not null,
  type          input_types           not null,
  required      boolean               not null,
  default_value varchar(255),
  options       jsonb,
  merchant      varchar(32),
  enabled       boolean default true  not null,
  searchable    boolean default false not null
);

comment on table attributes
is '属性表';

comment on column attributes.type
is '输入类型';

comment on column attributes.required
is '是否必须';

comment on column attributes.default_value
is '默认值';

comment on column attributes.options
is '选项值';

comment on column attributes.merchant
is '商家ID，公共的为空';

create table product_images
(
  id           varchar(32)  not null
    constraint product_images_pkey
    primary key,
  product_id   varchar(32)  not null,
  small_image  varchar(255) not null,
  medium_image varchar(255) not null,
  big_image    varchar(255) not null
);

comment on table product_images
is '商品图片表';

comment on column product_images.small_image
is '缩略图';

comment on column product_images.medium_image
is '中等大小图片';

comment on column product_images.big_image
is '大图（原图）';

create table merchant_profiles
(
  user_id          varchar(32) not null
    constraint merchant_profiles_pkey
    primary key,
  social_id        varchar(18),
  "rep_name
  
  
  "              varchar(32),
  company_name     varchar(255),
  company_address  address,
  delivery_address address,
  license_image    varchar(255),
  company_image    varchar(255),
  debit_card_info  debit_card
);

comment on table merchant_profiles
is '商家信息表';

comment on column merchant_profiles.social_id
is '身份证号';

comment on column merchant_profiles."rep_name


"
is '法人姓名/真实姓名';

comment on column merchant_profiles.company_name
is '公司名称';

comment on column merchant_profiles.company_address
is '公司地址';

comment on column merchant_profiles.delivery_address
is '发货地';

comment on column merchant_profiles.license_image
is '营业执照图片';

comment on column merchant_profiles.company_image
is '形象照';

create table attribute_sets
(
  id            varchar(32)  not null
    constraint attribute_sets_pkey
    primary key,
  name          varchar(255) not null,
  attribute_ids character varying []
);

comment on table attribute_sets
is '属性集表';

comment on column attribute_sets.attribute_ids
is '属性ID';

create table brands
(
  id          varchar(32)           not null
    constraint brands_pkey
    primary key,
  name        varchar(255),
  thumbnaim   varchar(255),
  description text,
  deleted     boolean default false not null,
  enabled     boolean default true  not null,
  sort_order  integer default 255   not null
);

comment on table brands
is '品牌表';

comment on column brands.thumbnaim
is '缩略图';

create table operation_logs
(
  id         varchar(32)                         not null
    constraint operation_logs_pkey
    primary key,
  user_id    varchar(32)                         not null,
  action     varchar(255)                        not null,
  created_at timestamp default CURRENT_TIMESTAMP not null
);

comment on table operation_logs
is '操作日志表';

create table comments
(
  id         varchar(32)                         not null
    constraint comments_pkey
    primary key,
  user_id    varchar(32)                         not null,
  product_id varchar(32)                         not null,
  rank       smallint,
  order_id   varchar(32)                         not null,
  content    text                                not null,
  is_show    boolean default false               not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  parent_id  varchar(32),
  user_ip    inet
);

comment on table comments
is '评价表';

comment on column comments.rank
is '星级;只有1 到5 星;由数字代替;其中5 代表5 星';

comment on column comments.is_show
is '是否审核通过（通过才显示）';

comment on column comments.user_ip
is '用户评论时IP';

create table payment
(
  id         varchar(32)           not null
    constraint payment_pkey
    primary key,
  pay_name   varchar(255)          not null,
  pay_code   varchar(32)           not null,
  pay_fee    text,
  pay_desc   text,
  sort_order integer default 255   not null,
  enabled    boolean default false not null,
  is_cod     boolean default false not null,
  is_online  boolean default false not null
);

comment on table payment
is '支付方式表';

comment on column payment.pay_code
is '支付方式 的英文缩写,其实是该支付方式处理插件的不带后缀的文件 名部分';

comment on column payment.pay_fee
is '支付费用';

comment on column payment.pay_desc
is '支付方式描述';

comment on column payment.sort_order
is '支付方式的显示顺序';

comment on column payment.enabled
is '是否可用';

comment on column payment.is_cod
is '是否货到付款';

comment on column payment.is_online
is '是否在线支付';

create table configs
(
  id         varchar(32)          not null
    constraint config_pkey
    primary key,
  parent_id  varchar(32),
  code       varchar(32)          not null,
  type       input_types,
  value      text                 not null,
  sort_order smallint default 255 not null
);

comment on table configs
is '全站配置表';

comment on column configs.code
is '跟变量名的作用差不多，其实就是语言包中的字符串索引，如$_LANG[''''cfg_range''''][''''cart_confirm'''']';

comment on column configs.type
is '该配置的类型';

create table order_products
(
  id             varchar(32)  not null
    constraint order_products_pkey
    primary key,
  order_id       varchar(32)  not null,
  product_id     varchar(32)  not null,
  product_name   varchar(255) not null,
  product_number smallint     not null,
  product_price  money        not null,
  product_image  varchar(255) not null
);

comment on table order_products
is '商品订单关联表';

comment on column order_products.product_name
is '下单时的商品名';

comment on column order_products.product_number
is '购买数量';

comment on column order_products.product_price
is '下单时商品的售价';

comment on column order_products.product_image
is '商品缩略图';

create table favorites
(
  "id "      varchar(32)                         not null,
  user_id    varchar(32)                         not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  type       favorite_type                       not null,
  object_id  varchar(32)                         not null,
  constraint favorites_pkey
  primary key (id)
);

comment on table favorites
is '收藏表';

comment on column favorites.type
is '收藏对象的类型';

create table cart
(
  id            varchar(32) not null
    constraint cart_pkey
    primary key,
  product_id    varchar(32) not null,
  user_id       varchar(32) not null,
  product_count integer     not null,
  merchant_id   varchar(32) not null
);

comment on table cart
is '购物车';

comment on column cart.product_count
is '商品数量';

comment on column cart.merchant_id
is '商家ID';

create table cash_request
(
  id              varchar(32)                                                    not null
    constraint cash_request_pkey
    primary key,
  amount          money                                                          not null,
  debit_card_info debit_card                                                     not null,
  status          cash_request_status default 'unchecked' :: cash_request_status not null,
  reply           text,
  note            text
);


