import React, {useState} from "react";
import {Alert, Col, Form} from "react-bootstrap";
import classNames from "classnames";
import YgoIcon from "../../core/YgoIcon";
import {DraftMode} from "../../api/Draft";

type SettingsModeSelectEntryProps = {
    value: DraftMode
    setValue: React.Dispatch<React.SetStateAction<DraftMode>>
    className?: string
    md: 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
};

function SettingsModeSelectEntry(props: SettingsModeSelectEntryProps) {
    const [expanded, setExpanded] = useState<boolean>(false)

    const CollapsedIcon = <YgoIcon icon={"help-outline"}
                                   size={18}
                                   onClick={(event) => {
                                       setExpanded(true)
                                       event.preventDefault()
                                   }}
                                   classNames={classNames("fill-neutral-900 dark:fill-neutral-50", "hover:fill-blue-500 hover:dark:fill-blue-200", "active:hover:fill-blue-600 active:hover:dark:fill-blue-300")}/>
    const ExpandIcon = <YgoIcon icon={"help-fill"}
                                size={18}
                                onClick={(event) => {
                                    setExpanded(false)
                                    event.preventDefault()
                                }}
                                classNames={classNames("fill-blue-600 dark:fill-blue-400", "hover:fill-blue-500 hover:dark:fill-blue-300", "active:hover:fill-blue-400 active:hover:dark:fill-blue-200")}/>

    const modeButtonsCN = classNames("p-2 border-top border-bottom dark:text-white")
    const modeButtonsSelectedCN = classNames("bg-blue-600 text-white")
    return <Form.Group as={Col} md={props.md} className={classNames(props.className ? props.className : "")}>
        <Form.Label className={"w-100"}>
            <div className={"flex"}>
                <div className={"self-center fw-bold mr-2 dark:text-white"}>Draft Mode</div>
                {expanded ? ExpandIcon : CollapsedIcon}
            </div>
            {expanded ? <div>
                <Alert className={"mb-0 mt-1"}>Defines the mode for the draft:<br/>
                    <span className={"fw-bold"}>Rounds</span> - Play a fixed amount of rounds.<br/>
                    <span className={"fw-bold"}>BestOf</span> - Play a game until one player wins at leat half of the
                    rounds.</Alert>
            </div> : <></>}
        </Form.Label>
        <div>
            <button className={classNames("border-start border-end rounded-tl rounded-bl", modeButtonsCN, props.value === "rounds" ? modeButtonsSelectedCN : "")} onClick={event => {
                props.setValue("rounds")
                event.preventDefault()
            }
            }>
                Rounds
            </button>
            <button className={classNames("border-end rounded-tr rounded-br", modeButtonsCN, props.value === "bestof" ? modeButtonsSelectedCN : "")} onClick={event => {
                props.setValue("bestof")
                event.preventDefault()
            }
            }>
                BestOf
            </button>
        </div>
    </Form.Group>
}

export default SettingsModeSelectEntry
