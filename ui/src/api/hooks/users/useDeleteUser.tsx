import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

type DeleteUserPayload = {
    email: string
}

const request: MutationFunction<string, string> = (email:string) => {
    const deleteUserData: DeleteUserPayload = {
        email: email
    }
    return axios.delete(`${PUBLIC_URL}/api/v1/users`, {
        data: deleteUserData
    })
}

export const useDeleteUser = (options?: Omit<UseMutationOptions<string, Error, string, unknown>, "mutationFn"> | undefined ) => {
    return useMutation<string, Error, string>(request, options);
};