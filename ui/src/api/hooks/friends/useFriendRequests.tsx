import {UseQueryResult} from "react-query";
import {GetFriendRequestsResponse} from "../../UserModel";
import {useMagicMethodAxios} from "../cards/useCards";

export function useFriendRequests(queryOptions: any = {}): UseQueryResult<GetFriendRequestsResponse> {
    return useMagicMethodAxios<GetFriendRequestsResponse>(["friendRequests"], `user/friends/requests`, new Map<string, string>(), queryOptions)
}