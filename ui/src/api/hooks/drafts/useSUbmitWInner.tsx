import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

export type SubmitWinnerRequest
    = {
    roundID: number
    winner: number
}

export type SubmitWinnerPayload = {
    winner: number
}

const request: MutationFunction<string, SubmitWinnerRequest
    > = (request) => {
    const payload:SubmitWinnerPayload = {
        winner: request.winner
    }
    return axios.post(`${PUBLIC_URL}/api/v1/rounds/${request.roundID}`, payload)
}

export const useSubmitWinner = (options?: Omit<UseMutationOptions<string, Error, SubmitWinnerRequest
    , unknown>, "mutationFn"> | undefined) => {
    return useMutation<string, Error, SubmitWinnerRequest
        >(request, options);
};