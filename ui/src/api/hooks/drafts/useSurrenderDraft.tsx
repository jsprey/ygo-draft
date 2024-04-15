import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

const request: MutationFunction<string, number
    > = (draftID) => {
    return axios.post(`${PUBLIC_URL}/api/v1/drafts/${draftID}/surrender`)
}

export const useSurrenderDraft = (options?: Omit<UseMutationOptions<string, Error, number>, "mutationFn"> | undefined) => {
    return useMutation<string, Error, number
        >(request, options);
};