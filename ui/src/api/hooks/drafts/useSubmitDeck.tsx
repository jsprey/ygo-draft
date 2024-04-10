import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";

export type SubmitDeckRequest = {
    roundID: number
    deck: string[]
}

export type SubmitDeckPayload = {
    deck: string[]
}

const request: MutationFunction<string, SubmitDeckRequest> = (request) => {
    const payload:SubmitDeckPayload = {
        deck: request.deck
    }
    return axios.post(`${PUBLIC_URL}/api/v1/rounds/${request.roundID}/decks`, payload)
}

export const useSubmitDeck = (options?: Omit<UseMutationOptions<string, Error, SubmitDeckRequest, unknown>, "mutationFn"> | undefined) => {
    return useMutation<string, Error, SubmitDeckRequest>(request, options);
};