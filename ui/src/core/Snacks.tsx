import {enqueueSnackbar} from "notistack";

export function ShowSuccessfulSnack(message: string) {
    enqueueSnackbar(message, {
        autoHideDuration: 6000,
        variant: "success"
    })
}

export function ShowErrorSnack(message: string) {
    enqueueSnackbar(message, {
        autoHideDuration: 6000,
        variant: "error"
    })
}