import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

const request: MutationFunction<string, string> = (email) => {
    return axios.post(`${PUBLIC_URL}/api/v1/user/friends/requests`,
        {friend_email: email})
}

export const useSendFriendRequest = (options?: Omit<UseMutationOptions<string, Error, string, unknown>, "mutationFn"> | undefined) => {
    return useMutation<string, Error, string>(request, options);
};