INSERT INTO ygodraft.public.draft_store (id, display_name, description, category, rarity, tags, costs)
VALUES ('draw_one_more', 'Draft One Card', 'You can draft another card for the next round.', 'Extra Draws', 'common', 'positive_effect,gain_card', 123)
ON CONFLICT (id)
    DO UPDATE
    SET display_name = 'Draft One Card',
        description = 'You can draft another card for the next round.',
        category = 'Extra Draws',
        rarity = 'common',
        tags = 'positive_effect,gain_card',
        costs = 123;