import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

const request: MutationFunction<string, number> = (friendID) => {
    return axios.post(`${PUBLIC_URL}/api/v1/user/friends/requests/${friendID}`)
}

export const useFriendsAcceptRequest = (options?: Omit<UseMutationOptions<string, Error, number, unknown>, "mutationFn"> | undefined ) => {
    return useMutation<string, Error, number>(request, options);
};