import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";
import {DraftRound} from "../../DraftRound";

export type useDraftRoundsPayload = {
    rounds: DraftRound[]
}

export function useDraftRounds(id:number, queryOptions: any = {}): UseQueryResult<useDraftRoundsPayload> {
    return useMagicMethodAxios<useDraftRoundsPayload>(["draft", id, "rounds"], `drafts/${id}/rounds`, new Map<string, string>(), queryOptions)
}