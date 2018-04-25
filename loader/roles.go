package loader

import (
    "gopkg.in/nicksrandall/dataloader.v5"
    "pinjihui.com/pinjihui/repository"
    "golang.org/x/net/context"
    "pinjihui.com/pinjihui/model"
    "fmt"
)

func newRolesLoader() dataloader.BatchFunc {
    return loadRolesBatch
}

func loadRolesBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
    var (
        n       = len(keys)
        results = make([]*dataloader.Result, n)
    )

    roles, err := ctx.Value("roleRepository").(*repository.RoleRepository).FindByUserIds(keys.Keys())
    for i, key := range keys {
        rs := make([]*model.Role, 0)
        for _, role := range roles {
            if key.String() == role.UserId {
                rs = append(rs, &role.Role)
            }
        }
        results[i] = &dataloader.Result{Data: rs, Error: err}
    }


    return results
}

func LoadRoles(ctx context.Context, key string) ([]*model.Role, error) {
    var roles []*model.Role

    ldr, err := extract(ctx, rolesLoaderKey)
    if err != nil {
        fmt.Errorf("Error in extract rolesLoaderKey : %v", err)
        return nil, err
    }

    data, err := ldr.Load(ctx, dataloader.StringKey(key))()
    if err != nil {
        fmt.Errorf("Error in extract Load : %v", err)
        return nil, err
    }
    roles, ok := data.([]*model.Role)
    if !ok {
        return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", roles, data)
    }

    return roles, nil
}