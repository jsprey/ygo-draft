export type RoundStatus = "preparation" | "fighting" | "finished";

export interface DraftRound {
    "id": number,
    "draft_id": number,
    "round_number": number,
    "status": RoundStatus,
    "winner_user_id": number
}

