import React, {useState} from "react";
import {useLocation} from "react-router-dom";
import {Alert} from "react-bootstrap";
import PageSettings from "../draft-local/PageSettings";
import {DraftSettings} from "../draft-local/DeckDraftWizard";
import {usePrompt} from "../api/hooks/usePromptBlocker";

export type ChallengeDraftPageProps = {}

export type ChallengeDraftState = {
    friendID: number,
    friendName: string
}

function ChallengeDraftPage() {
    usePrompt("You did not challenge your friend. Are you sure you want to leave?", true);

    const [setDraftSettings] = useState<DraftSettings>({} as DraftSettings)
    const location = useLocation();

    if (!location.state) {
        return <div className={"pt-2 pb-1"}>
            <Alert variant={'danger'}>Missing friend to invite. Try again.</Alert>
        </div>
    } else {
        const state = location.state as ChallengeDraftState;
        return <div className={"pb-2"}>
            <div className={"flex justify-content-center pt-4 pb-3"}>
                <p className={"text-5xl align-text-center uppercase"}>Challenge: {state.friendName}</p>
            </div>

            <PageSettings setDraftSettings={setDraftSettings} local={false} submitButtonName={"Challenge"} onSettingsSubmit={() => console.log("Challenge")}/>
        </div>
    }

}


export default ChallengeDraftPage
