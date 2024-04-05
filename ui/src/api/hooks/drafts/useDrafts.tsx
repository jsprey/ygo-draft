import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";
import {DraftChallenge} from "../../Draft";

export interface GetDraftsResponse {
    challenges: DraftChallenge[]
}

export function useDrafts(queryOptions: any = {}): UseQueryResult<GetDraftsResponse> {
    return useMagicMethodAxios<GetDraftsResponse>(["drafts", "pending"], `drafts/challenges`, new Map<string, string>(), queryOptions)
}