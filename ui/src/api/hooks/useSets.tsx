import {UseQueryResult} from "react-query";
import {useMagicMethod} from "./useCards";
import {SetList, SetWithCards} from "../Sets";

export function useSets(queryOptions: any = {}): UseQueryResult<SetList> {
    return useMagicMethod<SetList>(["sets"], `sets`, new Map<string, string>(), queryOptions)
}

export function useSetCards(setCode: string, queryOptions: any = {}): UseQueryResult<SetWithCards> {
    return useMagicMethod<SetWithCards>(["sets", setCode, "cards"], `sets/${setCode}/cards`, new Map<string, string>(), queryOptions)
}