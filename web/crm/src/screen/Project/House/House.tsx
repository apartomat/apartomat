import React, { useState } from "react"

import { Box, Button, BoxExtendedProps } from "grommet"

import { ProjectHouses } from "../useProject"

import { House as HouseType } from "./Add/useAddHouse"
import { Update } from "./Update/Update"
import { Add } from "./Add/Add"

export default function House({
    houses,
    projectId,
    onAdd,
    onUpdate,
    ...boxProps
}: {
    houses: ProjectHouses
    projectId: string
    onAdd?: (house: HouseType) => void
    onUpdate?: (house: HouseType) => void
} & BoxExtendedProps) {
    const [show, setShow] = useState(false)

    const [house, setHouse] = useState<HouseType | undefined>(first(houses))

    return (
        <Box {...boxProps}>
            <Box direction="row">
                <Button primary color="light-2" label={address(house) || "не указан"} onClick={() => setShow(!show)} />
            </Box>

            {show && house && (
                <Update
                    house={house}
                    onUpdate={(house) => {
                        setHouse(house)
                        setShow(false)
                        onUpdate && onUpdate(house)
                    }}
                    onEsc={() => setShow(false)}
                    onClickClose={() => setShow(false)}
                />
            )}

            {show && !house && (
                <Add
                    projectId={projectId}
                    onAdd={(house) => {
                        setHouse(house)
                        setShow(false)
                        onAdd && onAdd(house)
                    }}
                    onEsc={() => setShow(false)}
                    onClickClose={() => setShow(false)}
                />
            )}
        </Box>
    )
}

function first(houses: ProjectHouses) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            return houses.list.items[0]
    }

    return undefined
}

function address(house?: HouseType): string {
    if (house) {
        return [house.city, house.address, house.housingComplex].join(", ").replace(/,\s*$/, "")
    }

    return ""
}
