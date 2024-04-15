import {CardSet} from "./Sets";

export type DraftMode = "bestof" | "rounds" | undefined;
export type DraftStatus = "pending" | "running" | "declined" | "surrender" | "finished";

export interface DraftSettings {
    extra_deck_draws: number,
    extra_deck_size: number,
    main_deck_draws: number,
    main_deck_size: number
    mode: DraftMode,
    mode_value: number,
    sets: CardSet[]
}

export interface Draft {
    id: number,
    challenger_id: number,
    receiver_id: number,
    current_round_number: number,
    maximum_round_number: number,
    winner_user_id: number,
    challenge_date: Date,
    status: DraftStatus,
    settings: DraftSettings
}