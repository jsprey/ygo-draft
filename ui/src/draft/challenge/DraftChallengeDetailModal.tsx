import {Modal, Spinner} from "react-bootstrap";
import {DraftChallenge} from "../../api/Draft";
import {useAcceptChallenge} from "../../api/hooks/challenges/useAcceptChallenge";
import {useDeclineChallenge} from "../../api/hooks/challenges/useDeclineChallenge";
import {enqueueSnackbar} from "notistack";

export type DraftChallengeDetailModalProps = {
    challenge: DraftChallenge
    isShowing: boolean
    setShow: React.Dispatch<React.SetStateAction<boolean>>
}

function DraftChallengeDetailModal(props: DraftChallengeDetailModalProps) {
    const handleClose = () => props.setShow(false);

    function showSuccess(message: string) {
        enqueueSnackbar(message, {
            autoHideDuration: 6000,
            variant: "success"
        })

        props.setShow(false)
    }
    function showError(message: string) {
        enqueueSnackbar(message, {
            autoHideDuration: 6000,
            variant: "error"
        })
    }

    const acceptChallengeMutation = useAcceptChallenge({
        onSuccess: () => showSuccess("Challenge accepted."),
        onError: () => showError("Failed to accept challenge. Try again and/or contact the support.")
    });
    const declineChallengeMutation = useDeclineChallenge({
        onSuccess: () => showSuccess("Challenge declined."),
        onError: () => showError("Failed to decline challenge. Try again and/or contact the support.")
    });

    if (!props.isShowing) {
        return <></>
    }

    const settings = props.challenge.settings
    console.log(settings)
    return <Modal show={props.isShowing}
                  onHide={handleClose}
                  size={"xl"}
                  contentClassName={""}>
        <Modal.Header className={"bg-ygo-light dark:bg-ygo-dark border dark:text-white"}>
            <div className={"text-xl fw-bold"}>
                You got a challenge!
            </div>
        </Modal.Header>
        <Modal.Body className={"bg-ygo-light dark:bg-ygo-dark border dark:text-white"}>
            <div className={"fw-bold"}>Settings</div>
            <div className={"grid grid-cols-8"}>
                <span className={"col-span-2"}>Mode:</span>
                <span
                    className={"col-span-6"}>{settings.mode == "bestof" ? `Best of ${settings.modeValue} Rounds` : `${settings.modeValue} Rounds`}</span>
                <span className={"col-span-2"}>Main Deck Drafts:</span>
                <span className={"col-span-6"}>{settings.main_deck_draws}</span>
                <span className={"col-span-2"}>Main Deck Drafts Size:</span>
                <span className={"col-span-6"}>{settings.main_deck_size}</span>
                <span className={"col-span-2"}>Extra Deck Drafts:</span>
                <span className={"col-span-6"}>{settings.extra_deck_draws}</span>
                <span className={"col-span-2"}>Extra Deck Drafts Size:</span>
                <span className={"col-span-6"}>{settings.extra_deck_size}</span>
                <span className={"col-span-2"}>Sets:</span>
                <span className={"col-span-6"}>{settings.sets.map((value, index) => {
                        return (index == 0 ? "" : ", ") + value.set_name
                    }
                )}</span>
            </div>


        </Modal.Body>
        <Modal.Footer className={"bg-ygo-light dark:bg-ygo-dark border dark:text-white"}>
            <button disabled={declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading} className={"btn btn-success"} onClick={() => {
                acceptChallengeMutation.mutate(props.challenge.id)
            }
            }>
                {declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading ? <Spinner animation={"border"} size={"sm"}></Spinner> : "Accept"}
            </button>
            <button disabled={declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading} className={"btn btn-danger"} onClick={() => {
                declineChallengeMutation.mutate(props.challenge.id)
            }
            }>
                {declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading ? <Spinner animation={"border"} size={"sm"}></Spinner> : "Decline"}
            </button>
        </Modal.Footer>
    </Modal>
}

export default DraftChallengeDetailModal