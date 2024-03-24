import {UseQueryResult} from "react-query";
import {GetFriendsResponse} from "../../UserModel";
import {useMagicMethodAxios} from "../useCards";

export function useFriends(queryOptions: any = {}): UseQueryResult<GetFriendsResponse> {
    return useMagicMethodAxios<GetFriendsResponse>(["friends"], `user/friends`, new Map<string, string>(), queryOptions)
}