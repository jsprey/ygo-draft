import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

export type AddUserPayload = {
    email: string
    password: string
    display_name: string
    is_admin: boolean
}

const request: MutationFunction<string, AddUserPayload> = (payload: AddUserPayload) => {
    return axios.post(`${PUBLIC_URL}/api/v1/users`, payload)
}

export const useAddUser = (options?: Omit<UseMutationOptions<string, Error, AddUserPayload, unknown>, "mutationFn"> | undefined ) => {
    return useMutation<string, Error, AddUserPayload>(request, options);
};