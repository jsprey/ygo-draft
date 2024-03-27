import {UseQueryResult} from "react-query";
import {GetUserReponse} from "../../UserModel";
import {useMagicMethodAxios} from "../cards/useCards";

export function useCurrentUser(queryOptions: any = {}): UseQueryResult<GetUserReponse> {
    return useMagicMethodAxios<GetUserReponse>(["user_me"], `user`, new Map<string, string>(), queryOptions)
}