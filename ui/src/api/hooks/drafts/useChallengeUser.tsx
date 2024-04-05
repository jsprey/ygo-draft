import {MutationFunction, useMutation, UseMutationOptions} from "react-query";
import axios from "axios";
import {PUBLIC_URL} from "../../../index";
import {DraftSettings} from "../../Draft";

export interface PostDraftChallengeRequest {
    friend_id: number,
    settings: DraftSettings
}

const request: MutationFunction<string, PostDraftChallengeRequest> = (request) => {
    return axios.post(`${PUBLIC_URL}/api/v1/drafts`, request)
}

export const useChallengeUser = (options?: Omit<UseMutationOptions<string, Error, PostDraftChallengeRequest, unknown>, "mutationFn"> | undefined) => {
    return useMutation<string, Error, PostDraftChallengeRequest>(request, options);
};