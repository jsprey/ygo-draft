import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";
import {Draft} from "../../Draft";

export interface GetDraftsResponse {
    drafts: Draft[]
}

export function useDrafts(queryOptions: any = {}): UseQueryResult<GetDraftsResponse> {
    return useMagicMethodAxios<GetDraftsResponse>(["drafts", "running"], `drafts`, new Map<string, string>(), queryOptions)
}