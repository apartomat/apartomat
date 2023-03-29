import React, { ChangeEventHandler, useState } from "react"

import { Form, FormField, MaskedInput, TextInput } from "grommet"

export type Value = { name: string, square: string, level: string }

const defaultSuggestions = [
    "Ванная",
    "Гардероб",
    "Гостиная",
    "Детская",
    "Душевая",
    "Кабинет",
    "Кухня-гостиная",
    "Лестница",
    "Прихожая",
    "Санузел",
    "Спальня",
    "Туалет",
];

const filterSuggestions = (vals: string[], val: string) => {
    const re = new RegExp(`${val}`, 'i')

    return val ? vals.filter(v => re.test(v) && v !== val) : vals
} 

export default function Room({
    value,
    onChange,
    onSubmit,
    submit
}: {
    value: Value,
    onChange?: (value: Value) => void,
    onSubmit?: (event: React.FormEvent) => void,
    submit?: JSX.Element
}) {
    const [ suggestions, setSuggestions ] = useState<string[]>(filterSuggestions(defaultSuggestions, value.name))

    const handleChangeName: ChangeEventHandler<HTMLInputElement> = ({ target: { value }}) => {
        setSuggestions(filterSuggestions(defaultSuggestions, value))
    }

    return (
        <Form
            validate="submit"
            value={value}
            onChange={val => onChange && onChange(val)}
            onSubmit={onSubmit}
            messages={{required: "обязательное поле"}}
        >
            <FormField
                name="name"
                label="Название"
                width="medium"
                onChange={handleChangeName}
            >
                <TextInput
                    name="name"
                    suggestions={suggestions}
                    placeholder="Комната"
                />
            </FormField>
            <FormField
                name="square"
                label="Площадь,&nbsp;м²"
                width="xsmall"
            >
                <MaskedInput
                    name="square"
                    mask={[
                        { regexp: /^[\d*,\d*]+$/, placeholder: "0,0" },
                    ]}
                    width="xsmall"
                />
            </FormField>
            {submit}
        </Form>
    )
}