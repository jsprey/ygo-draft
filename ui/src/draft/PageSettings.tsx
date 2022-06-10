import React, {useState} from "react";
import {DraftSettings, DraftStages} from "./DeckDraftWizard";
import {Button, Form, Row} from "react-bootstrap";
import SettingsEntry from "./SettingsEntry";
import CardSetSelector from "./CardSetSelector";
import {CardSet} from "../api/Sets";

export type PageSettingsProps = {
    setCurrentStage: React.Dispatch<React.SetStateAction<DraftStages>>
    setDraftSettings: React.Dispatch<React.SetStateAction<DraftSettings>>
}

function PageSettings(props: PageSettingsProps) {
    const [generalDraftCardSets, setGeneralDraftCardSets] = useState<CardSet[]>([])
    const [mainDraftRound, setMainDraftRound] = useState(5)
    const [mainDraftRoundError, setMainDraftRoundError] = useState("")
    const [mainDraftSize, setMainDraftSize] = useState(3)
    const [mainDraftSizeError, setMainDraftSizeError] = useState("");
    const [extraDraftRound, setExtraDraftRound] = useState(5)
    const [extraDraftRoundError, setExtraDraftRoundError] = useState("")
    const [extraDraftSize, setExtraDraftSize] = useState(2)
    const [extraDraftSizeError, setExtraDraftSizeError] = useState("")
    const [validated, setValidated] = useState(false)

    const handleSubmit = (event: any) => {
        event.preventDefault();
        event.stopPropagation();

        if (mainDraftSizeError === "" && mainDraftRoundError === "" && extraDraftRoundError === "" && extraDraftSizeError === "") {
            setValidated(true)

            const draftSettings: DraftSettings = {
                mainDraftSize: mainDraftSize,
                mainDraftRound: mainDraftRound,
                extraDraftRound: extraDraftRound,
                extraDraftSize: extraDraftSize,
                selectedCardSets: generalDraftCardSets
            }
            props.setDraftSettings(draftSettings)
            props.setCurrentStage(DraftStages.DraftMain)
        }
    };

    return <>
        <div className={"p-3 mt-0 bg-blue-200 rounded-2 shadow-md mb-2"}>
            At this page it is possible to configure multiple aspects of the drafting phase. Keep in mind to use the
            same options as your dueling partner.
        </div>

        <Form className={"mt-3"} onSubmit={handleSubmit} noValidate validated={validated}>
            <div className={"title  mt-0 fw-bold"}>General Settings</div>

            <CardSetSelector selectedSets={generalDraftCardSets} setSelectedSets={setGeneralDraftCardSets} rowClass={"mt-0"}
                             tooltip={"Only cards from the defined sets are used when drafting a deck."}></CardSetSelector>
            <div className={"title mt-4 fw-bold"}>Settings for the Main Draft</div>
            <Row>
                <SettingsEntry value={mainDraftRound} setValue={setMainDraftRound} md={6} error={mainDraftRoundError}
                               setError={setMainDraftRoundError} min={5} max={80}
                               title={"Number of Rounds (Main Draft)"}
                               tooltip={"Defines the number of rounds while drafting the main deck. The resulting main deck will have the same size as the round number. Valid Values: [40-80]."}/>

                <SettingsEntry value={mainDraftSize} setValue={setMainDraftSize} md={6} error={mainDraftSizeError}
                               setError={setMainDraftSizeError} min={2} max={10}
                               title={"Card Each Round (Main Draft)"}
                               tooltip={"Defines the number of cards that are proposed for every round of the draft while drafting the main deck.. Valid Values: [2-10]."}/>
            </Row>
            <div className={"title mt-4 fw-bold"}>Settings for the Extra Draft</div>
            <Row>
                <SettingsEntry value={extraDraftRound} setValue={setExtraDraftRound} md={6} error={extraDraftRoundError}
                               setError={setExtraDraftRoundError} min={0} max={20}
                               title={"Number of Rounds (Extra Draft)"}
                               tooltip={"Defines the number of rounds while drafting the extra deck. The resulting extra deck will have the same size as the round number. Valid Values: [0-20]."}/>

                <SettingsEntry value={extraDraftSize} setValue={setExtraDraftSize} md={6} error={extraDraftSizeError}
                               setError={setExtraDraftSizeError} min={2} max={10}
                               title={"Card Each Round (Extra Draft)"}
                               tooltip={"Defines the number of cards that are proposed for every round of the draft while drafting the extra deck. Valid Values: [2-10]."}/>
            </Row>
            <Button className={"mt-3"} type="submit"
                    disabled={mainDraftSizeError !== "" || mainDraftRoundError !== "" || extraDraftSizeError !== "" || extraDraftRoundError !== ""}>Next</Button>
        </Form>
    </>
}


export default PageSettings
