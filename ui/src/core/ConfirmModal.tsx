import React, {Dispatch, Fragment, useRef} from "react";
import { Dialog, Transition } from '@headlessui/react'
import Button from "./Button";
import YgoIcon from "./YgoIcon";

type ConfirmModalProps = {
    show: boolean
    setShow: Dispatch<React.SetStateAction<boolean>>
    onConfirm?: () => void
    onCancel?: () => void
    title: string
    description: string
    confirmName: string
};

function ConfirmModal(props: ConfirmModalProps) {
    let cancelButtonRef = useRef(null)

    return <Transition.Root show={props.show} as={Fragment}>
        <Dialog as="div"
                initialFocus={cancelButtonRef}
                className="relative z-10 text-dark dark:text-light"
                onClose={props.setShow}>
            <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0"
                enterTo="opacity-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
            >
                <div className="fixed inset-0 bg-dark-2 bg-opacity-75 transition-opacity" />
            </Transition.Child>

            <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
                <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                        enterTo="opacity-100 translate-y-0 sm:scale-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                        leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                    >
                        <Dialog.Panel className="relative transform overflow-hidden rounded-lg text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                            <div className="bg-light dark:bg-dark px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                                <div className="sm:flex sm:items-start">
                                    <div className="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-danger-light sm:mx-0 sm:h-10 sm:w-10">
                                        <YgoIcon icon={"xcircle"} size={18} classNames={"fill-danger-dark"}/>
                                    </div>
                                    <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                                        <Dialog.Title as="h3" className="text-base font-semibold leading-6 text-dark dark:text-light">
                                            {props.title}
                                        </Dialog.Title>
                                        <div className="mt-2">
                                            <p className="text-sm text-dark-1 dark:text-light-1">
                                                {props.description}
                                            </p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="bg-light-1 dark:bg-dark-1 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                                <Button
                                    variant={"danger"}
                                    className="inline-flex w-full justify-center px-3 py-2 font-semibold sm:ml-3 sm:w-auto"
                                    onClick={() => {
                                        props.setShow(false)
                                        if (props.onConfirm) {
                                            props.onConfirm()
                                        }
                                    }}
                                >
                                    {props.confirmName}
                                </Button>
                                <Button
                                    variant={"neutral"}
                                    className="mt-3 inline-flex w-full justify-center px-3 py-2 text-sm font-semibold sm:mt-0 sm:w-auto"
                                    ref={cancelButtonRef}
                                    onClick={() => {
                                        props.setShow(false)
                                        if (props.onCancel) {
                                            props.onCancel()
                                        }
                                    }}
                                >
                                    Cancel
                                </Button>
                            </div>
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </div>
        </Dialog>
    </Transition.Root>
}

export default ConfirmModal
