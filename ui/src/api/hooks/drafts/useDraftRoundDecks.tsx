import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";

export type useDraftRoundDecksPayload = {
    user_deck: string[]
    enemy_deck: string[]
}

export function useDraftRoundDecks(roundID:string, queryOptions: any = {}): UseQueryResult<useDraftRoundDecksPayload> {
    return useMagicMethodAxios<useDraftRoundDecksPayload>(["rounds", roundID, "decks"], `rounds/${roundID}/decks`, new Map<string, string>(), queryOptions)
}