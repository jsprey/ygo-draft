import React, {Dispatch, Fragment, ReactNode} from "react";
import {Dialog, Transition} from '@headlessui/react'

type ModalProps = {
    show: boolean
    setShow: Dispatch<React.SetStateAction<boolean>>
    onHide?: () => void
    children: ReactNode | undefined
};

function Modal(props: ModalProps) {
    return <Transition.Root show={props.show} as={Fragment}>
        <Dialog as="div"
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
                        <Dialog.Panel style={{maxWidth: "80%", maxHeight: "80%"}} className="text-dark dark:text-light bg-light dark:bg-dark p-2 relative transform overflow-hidden rounded-lg text-left shadow-xl transition-all sm:my-8 sm:w-full border border-border">
                            {props.children ? props.children : null}
                        </Dialog.Panel>
                    </Transition.Child>
                </div>
            </div>
        </Dialog>
    </Transition.Root>
}

export default Modal
