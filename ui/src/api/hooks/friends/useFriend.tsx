import {UseQueryResult} from "react-query";
import {Friend} from "../../UserModel";
import {useMagicMethodAxios} from "../cards/useCards";

export function useFriend(id: number, queryOptions: any = {}): UseQueryResult<Friend> {
    return useMagicMethodAxios<Friend>(["friends", id], `user/friends/${id}`, new Map<string, string>(), queryOptions)
}