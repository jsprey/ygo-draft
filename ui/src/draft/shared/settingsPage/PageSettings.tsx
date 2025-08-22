import React from "react";
import SettingsEntry from "./SettingsEntry";
import CardSetSelector from "./CardSetSelector";
import {CardSet} from "../../../api/Sets";
import {DraftMode, DraftSettings} from "../../../api/Draft";
import classNames from "classnames";
import SettingsModeSelectEntry from "./SettingsModeSelectEntry";
import Alert from "../../../core/Alert";
import Button from "../../../core/Button";
import useLocalState from "../../../api/hooks/storage/useLocalState";

export type PageSettingsProps = {
    onSettingsSubmit: (settings: DraftSettings) => void
    submitButtonName: string
    local?: boolean
}

function PageSettings(props: PageSettingsProps) {
    const [generalDraftCardSets, setGeneralDraftCardSets, resetGeneralDraftCardSets] = useLocalState<CardSet[]>("page_settings_card_sets", [] as CardSet[])
    const [mode, setMode, resetMode] = useLocalState<DraftMode>("page_settings_mode", "rounds")
    const [numberOfRounds, setNumberOfRounds, resetNumberOfRounds] = useLocalState<number>("page_settings_number_of_rounds", 10)
    const [numberOfRoundsError, setNumberOfRoundsError, resetNumberOfRoundsError] = useLocalState<string>("page_settings_number_of_rounds_error", "")
    const [mainDraftRound, setMainDraftRound, resetMainDraftRound] = useLocalState<number>("page_settings_main_rounds", 5)
    const [mainDraftRoundError, setMainDraftRoundError, resetMainDraftRoundError] = useLocalState<string>("page_settings_main_rounds_error", "")
    const [mainDraftSize, setMainDraftSize, resetMainDraftSize] = useLocalState<number>("page_settings_main_size", 3)
    const [mainDraftSizeError, setMainDraftSizeError, resetMainDraftSizeError] = useLocalState<string>("page_settings_main_size_error", "");
    const [extraDraftRound, setExtraDraftRound, resetExtraDraftRound] = useLocalState<number>("page_settings_extra_rounds", 5)
    const [extraDraftRoundError, setExtraDraftRoundError, resetExtraDraftRoundError] = useLocalState<string>("page_settings_extra_rounds_error", "")
    const [extraDraftSize, setExtraDraftSize, resetExtraDraftSize] = useLocalState<number>("page_settings_extra_size", 2)
    const [extraDraftSizeError, setExtraDraftSizeError, resetExtraDraftSizeError] = useLocalState<string>("page_settings_extra_size_error", "")

    function resetStorage() {
        resetGeneralDraftCardSets()
        resetMode()
        resetNumberOfRounds()
        resetNumberOfRoundsError()
        resetMainDraftRound()
        resetMainDraftRoundError()
        resetMainDraftSize()
        resetMainDraftSizeError()
        resetExtraDraftRound()
        resetExtraDraftRoundError()
        resetExtraDraftSize()
        resetExtraDraftSizeError()
    }

    const isLocalSettings = props.local ? props.local : false

    function isValid(): boolean {
        return mode !== undefined && numberOfRoundsError === "" && mainDraftRoundError === "" && mainDraftSizeError === "" && extraDraftRoundError === "" && extraDraftSizeError === "";
    }

    const handleSubmit = () => {
        if (isValid()) {

            const draftSettings: DraftSettings = {
                mode: mode,
                mode_value: numberOfRounds,
                main_deck_size: mainDraftSize,
                main_deck_draws: mainDraftRound,
                extra_deck_draws: extraDraftRound,
                extra_deck_size: extraDraftSize,
                sets: generalDraftCardSets
            }

            props.onSettingsSubmit(draftSettings)
        }
    };

    function showDraftInformation() {
        return isLocalSettings ? <Alert variant={"info"} className={"mt-2 mb-2"}>
            {isLocalSettings ? "At this page it is possible to configure multiple aspects of the drafting phase. Keep in " +
                "mind to use the same options as your dueling partner." : "Define the settings of your challenge!"
            }
        </Alert> : null
    }

    const headingCN = classNames("title mb-2 text-xl border-b border-border font-bold dark:text-white")
    return <>
        {showDraftInformation()}

        <div className={"mt-3"} onSubmit={handleSubmit}>
            <div className={classNames(headingCN, "mt-0")}>General</div>
            <div className={"flex"}>
                <div className={"w-2/4"}>
                    <SettingsModeSelectEntry value={mode} setValue={setMode} className={"mb-2 pr-1"} md={6}/>
                </div>
                <div className={"w-2/4"}>
                    <SettingsEntry value={numberOfRounds}
                                   setValue={setNumberOfRounds}
                                   error={numberOfRoundsError}
                                   setError={setNumberOfRoundsError}
                                   min={2} max={20}
                                   label={"Number of Rounds"}
                                   className={"mb-2 pl-1"}
                                   tooltip={"Defines the number of bestof/normal rounds."}/>
                </div>
            </div>

            <CardSetSelector selectedSets={generalDraftCardSets} setSelectedSets={setGeneralDraftCardSets}
                             rowClass={"mt-2 w-full"}
                             tooltip={"Only cards from the defined sets are used when drafting a deck."}/>
            <div className={classNames(headingCN, "mt-4")}>Main Draft</div>
            <div className={"flex"}>
                <div className={"w-2/4"}>
                    <SettingsEntry value={mainDraftRound}
                                   setValue={setMainDraftRound}
                                   min={5} max={80}
                                   error={mainDraftRoundError}
                                   setError={setMainDraftRoundError}
                                   label={"Number of Draws"}
                                   className={"pr-1"}
                                   tooltip={"Defines the number of draw rounds while drafting the main deck. The resulting main deck will have the same size as the draw number. Valid Values: [40-80]."}/>
                </div>
                <div className={"w-2/4"}>
                    <SettingsEntry value={mainDraftSize} setValue={setMainDraftSize}
                                   min={2} max={10}
                                   error={mainDraftSizeError}
                                   setError={setMainDraftSizeError}
                                   label={"Card Each Draw"}
                                   className={"pl-1"}
                                   tooltip={"Defines the number of cards that are proposed for every draw round of the draft while drafting the main deck. Valid Values: [2-10]."}/>
                </div>

            </div>
            <div className={classNames(headingCN, "mt-4")}>Extra Draft</div>
            <div className={"flex"}>
                <div className={"w-2/4"}>
                    <SettingsEntry value={extraDraftRound}
                                   setValue={setExtraDraftRound}
                                   error={extraDraftRoundError}
                                   setError={setExtraDraftRoundError}
                                   min={0} max={20}
                                   className={"pr-1"}
                                   label={"Number of Draws"}
                                   tooltip={"Defines the number of draw rounds while drafting the extra deck. The resulting extra deck will have the same size as the draw number. Valid Values: [0-20]."}/>
                </div>
                <div className={"w-2/4"}>
                    <SettingsEntry value={extraDraftSize}
                                   setValue={setExtraDraftSize}
                                   error={extraDraftSizeError}
                                   setError={setExtraDraftSizeError}
                                   min={2} max={10}
                                   className={"pl-1"}
                                   label={"Card Each Draw"}
                                   tooltip={"Defines the number of cards that are proposed for every draw of the draft while drafting the extra deck. Valid Values: [2-10]."}/>
                </div>
            </div>
            <div className={"flex place-content-end"}>
                <Button variant={"secondary"} className={"mt-3 mr-2"} onClick={resetStorage}>
                    Default Values
                </Button>
                <Button variant={"primary"} className={"mt-3"} onClick={handleSubmit}
                        disabled={!isValid()}>
                    {props.submitButtonName}
                </Button>
            </div>
        </div>
    </>
}

export default PageSettings
