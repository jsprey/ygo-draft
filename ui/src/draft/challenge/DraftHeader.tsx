import React, {useState} from "react";
import {Draft} from "../../api/Draft";
import DraftSettingsDetails from "./DraftSettingsDetails";
import SvgIconButton from "../../core/SvgIconButton";
import classNames from "classnames";
import DraftBattleBadge from "./DraftBattleBadge";
import {User} from "../../api/hooks/users/useUsers";
import {Friend} from "../../api/UserModel";

const CollapsedIcon = <SvgIconButton size={18}
                                     classNames={"fill-gray-400 hover:fill-gray-300 active:fill-gray-200 dark:fill-gray-100 hover:dark:fill-gray-200 active:dark:fill-gray-300"}>
    <path
        d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z"/>
</SvgIconButton>
const ExpandIcon = <SvgIconButton size={18}
                                  classNames={"fill-gray-700 hover:fill-gray-600 active:fill-gray-500 dark:fill-gray-400 hover:dark:fill-gray-300 active:dark:fill-gray-200"}>
    <path
        d="M7.247 11.14 2.451 5.658C1.885 5.013 2.345 4 3.204 4h9.592a1 1 0 0 1 .753 1.659l-4.796 5.48a1 1 0 0 1-1.506 0z"/>
</SvgIconButton>

type DraftHeaderProps = {
    draft: Draft;
    player: User,
    enemy: Friend
};

function DraftHeader(props: DraftHeaderProps) {
    const [settingCollapsed, setSettingCollapsed] = useState<boolean>(true)

    const settingsHeaderCN = classNames("flex-grow-1 pl-2", "rounded-tl rounded-tr", settingCollapsed ? "rounded-bl rounded-br" : "", "border dark:border-white", "bg-gray-200 dark:bg-gray-600")
    const settingsBodyCN = classNames("p-2", "rounded-bl rounded-br", "border-start border-bottom border-end dark:border-white", "bg-gray-200 dark:bg-gray-600")

    return <div className={"w-100 text-gray-400 dark:text-gray-100"}>
        <DraftBattleBadge playerNameOne={props.player.display_name} playerNameTwo={props.enemy.name}/>

        <div className={classNames("flex justify-content-center", "mt-2", settingsHeaderCN)}>
            <div
                className={classNames("align-self-center mr-1 uppercase fw-bold")}>
                Draft Settings
            </div>
            <div className={"align-self-center"} onClick={() => setSettingCollapsed(!settingCollapsed)}>
                {settingCollapsed ? CollapsedIcon : ExpandIcon}
            </div>
        </div>
        {settingCollapsed ? <></> : <DraftSettingsDetails settings={props.draft.settings} containerClasses={settingsBodyCN}/>}
    </div>
}

export default DraftHeader
