import React from "react"

import { Form as FormElement, FormField, MaskedInput } from "grommet"

export type ContactFormData = { fullName: string; phone: string; email: string; instagram: string }

export function Form({
    contact,
    onSet,
    onSubmit,
    submit,
}: {
    contact: ContactFormData
    onSet: React.Dispatch<ContactFormData>
    onSubmit: (event: React.FormEvent) => void
    submit: JSX.Element
}) {
    return (
        <FormElement
            validate="submit"
            value={contact}
            onChange={(val) => onSet(val)}
            onSubmit={onSubmit}
            messages={{ required: "обязательное поле" }}
        >
            <FormField
                label="Имя"
                name="fullName"
                validate={{
                    regexp: /^.+$/,
                    message: "обязательно для заполнения",
                    status: "error",
                }}
            >
                <MaskedInput
                    name="fullName"
                    mask={[
                        { regexp: /^.*$/, placeholder: "Имя" },
                        { fixed: " " },
                        { regexp: /^.*$/, placeholder: "Фамилия" },
                    ]}
                />
            </FormField>
            <FormField label="Телефон" name="phone">
                <MaskedInput
                    name="phone"
                    mask={[
                        { fixed: "+7 (" },
                        {
                            length: 3,
                            regexp: /^[0-9]{1,3}$/,
                            placeholder: "xxx",
                        },
                        { fixed: ")" },
                        { fixed: " " },
                        {
                            length: 3,
                            regexp: /^[0-9]{1,3}$/,
                            placeholder: "xxx",
                        },
                        { fixed: "-" },
                        {
                            length: 2,
                            regexp: /^[0-9]{1,4}$/,
                            placeholder: "xx",
                        },
                        { fixed: "-" },
                        {
                            length: 2,
                            regexp: /^[0-9]{1,4}$/,
                            placeholder: "xx",
                        },
                    ]}
                />
            </FormField>
            <FormField label="Электронная почта" name="email">
                <MaskedInput
                    name="email"
                    mask={[
                        { regexp: /^[\w\-_.]+$/, placeholder: "example" },
                        { fixed: "@" },
                        { regexp: /^[\w\-_.]+$/, placeholder: "test" },
                        { fixed: "." },
                        { regexp: /^[\w]+$/, placeholder: "org" },
                    ]}
                />
            </FormField>
            <FormField label="Instagram" name="instagram">
                <MaskedInput name="instagram" mask={[{ fixed: "https://www.instagram.com/" }, { regexp: /^.*$/ }]} />
            </FormField>
            {submit}
        </FormElement>
    )
}

export default Form
