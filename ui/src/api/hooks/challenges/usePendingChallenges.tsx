import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";
import {DraftChallenge} from "../../Draft";

export interface GetPendingChallengesResponse {
    challenges: DraftChallenge[]
}

export function usePendingChallenges(queryOptions: any = {}): UseQueryResult<GetPendingChallengesResponse> {
    return useMagicMethodAxios<GetPendingChallengesResponse>(["challenges", "pending"], `drafts/challenges`, new Map<string, string>(), queryOptions)
}