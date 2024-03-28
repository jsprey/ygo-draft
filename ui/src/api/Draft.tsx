import {CardSet} from "./Sets";

export type DraftMode = "bestof" | "rounds";
export type DraftChallengeStatus = "pending" | "accepted" | "declined";

export interface DraftSettings {
    extra_deck_draws: number,
    extra_deck_size: number,
    main_deck_draws: number,
    main_deck_size: number
    mode: DraftMode,
    modeValue: number,
    sets: CardSet[]
}

export interface DraftChallenge {
    id: number,
    challenger_id: number,
    receiver_id: number,
    challenge_date: Date,
    status: DraftChallengeStatus,
    settings: DraftSettings
}