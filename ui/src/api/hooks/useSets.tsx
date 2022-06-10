import {UseQueryResult} from "react-query";
import {useMagicMethod} from "./useCards";
import {SetList} from "../Sets";

export function useSets(queryOptions: any = {}): UseQueryResult<SetList> {
    return useMagicMethod<SetList>(["sets"], `sets`, new Map<string, string>(), queryOptions)
}