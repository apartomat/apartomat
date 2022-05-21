import React from "react"

import { Box, Heading, Text } from "grommet"

import { ProjectHouses } from "../useProject"

export default function HouseText({ houses }: { houses: ProjectHouses }) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            const [ house ] =  houses.list.items

            if (!house) {
                return (
                    <>n/a</>
                )
            }

            return (
                <>{[house.city, house.address, house.housingComplex].join(', ')}</>
            )
        default:
            return (
                <>n/a</>
            )
    }
}

function House({ houses }: { houses: ProjectHouses }) {
    return (
        <Box margin="none">
            <Heading level={4} margin={{ bottom: "xsmall"}}>Адрес</Heading>
            <Box height="36px" justify="center">
                <Text><HouseText houses={houses}/></Text>
            </Box>
        </Box>
    )
}