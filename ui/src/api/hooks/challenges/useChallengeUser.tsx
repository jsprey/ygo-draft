import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";
import {DraftSettings} from "../../Draft";

export interface PostChallengeRequest {
    friend_id: number,
    settings: DraftSettings
}

const request: MutationFunction<string, PostChallengeRequest> = (request) => {
    return axios.post(`${PUBLIC_URL}/api/v1/drafts/challenges`, request)
}

export const useChallengeUser = (options?: Omit<UseMutationOptions<string, Error, PostChallengeRequest, unknown>, "mutationFn"> | undefined) => {
    return useMutation<string, Error, PostChallengeRequest>(request, options);
};