INSERT INTO ygodraft.public.draft_store (id, display_name, description, category, rarity, tags, costs)
VALUES ({{.ID}}, {{.DisplayName}}, {{.Description}}, {{.Category}}, {{.Rarity}}, {{.Tags}}, {{.Costs}})
ON CONFLICT (id)
    DO UPDATE
    SET display_name = {{.DisplayName}},
        description = {{.Description}},
        category = {{.Category}},
        rarity = {{.Rarity}},
        tags = {{.Tags}},
        costs = {{.Costs}};