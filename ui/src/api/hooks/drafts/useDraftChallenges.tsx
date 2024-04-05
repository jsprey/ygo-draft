import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";
import {Draft} from "../../Draft";

export interface GetDraftChallengesResponse {
    drafts: Draft[]
}

export function useDraftChallenges(queryOptions: any = {}): UseQueryResult<GetDraftChallengesResponse> {
    return useMagicMethodAxios<GetDraftChallengesResponse>(["drafts", "challenges"], `drafts/challenges`, new Map<string, string>(), queryOptions)
}