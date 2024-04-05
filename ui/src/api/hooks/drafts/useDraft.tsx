import {UseQueryResult} from "react-query";
import {useMagicMethodAxios} from "../cards/useCards";
import {Draft} from "../../Draft";

export function useDraft(id:string, queryOptions: any = {}): UseQueryResult<Draft> {
    return useMagicMethodAxios<Draft>(["draft", id], `drafts/${id}`, new Map<string, string>(), queryOptions)
}