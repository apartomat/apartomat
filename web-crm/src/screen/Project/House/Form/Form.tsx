import React from "react"
import {Form as FormElement, FormField, TextInput} from "grommet"

export type Value = { city: string, address: string, housingComplex: string }

export function Form({
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
    return (
        <FormElement
            validate="submit"
            value={value}
            onChange={val => onChange && onChange(val)}
            onSubmit={onSubmit}
            messages={{required: "обязательное поле"}}
        >
            <FormField
                label="Город"
                name="city"
                validate={{
                    regexp: /^.+$/,
                    message: "обязательно для заполнения",
                    status: "error"
                }}
            >
                <TextInput name="city"/>
            </FormField>
            <FormField label="Адрес" name="address">
                <TextInput name="address"/>
            </FormField>
            <FormField label="Жилой комлекс" name="housingComplex">
                <TextInput name="housingComplex"/>
            </FormField>
            {submit}
        </FormElement>
    )
}