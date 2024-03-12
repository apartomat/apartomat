import {
    Accordion,
    AccordionPanel,
    Box,
    BoxExtendedProps,
    Button,
    CheckBox,
    Form,
    FormField,
    Heading,
    Layer,
    LayerExtendedProps,
    Text,
    TextInput,
} from "grommet"
import { FormClose, Link } from "grommet-icons"

import { ProjectScreenPublicSiteFragment, PublicSiteStatus } from "api/graphql"
import React, { useEffect, useState } from "react"
import { useMakeProjectNotPublic, useMakeProjectPublic } from "pages/Project/PublicSite/useMakeProjectPublic"

export default function PublicSite({
    projectId,
    site,
    onChange,
    onClose,
    ...props
}: {
    projectId: string
    site: ProjectScreenPublicSiteFragment
    onChange?: () => void
    onClose?: (changed: boolean) => void
} & BoxExtendedProps) {
    const [showForm, setShowForm] = useState(false)

    return (
        <Box {...props} justify="center">
            <Box>
                <Button
                    label="Ссылка"
                    icon={<Link color={siteIsPublic(site) ? "status-ok" : "status-unknown"} />}
                    onClick={() => setShowForm(!showForm)}
                    title="Ссылка"
                />
            </Box>

            {showForm && (
                <EditForm
                    projectId={projectId}
                    publicSite={site}
                    onEsc={() => setShowForm(false)}
                    onClickClose={(changed: boolean) => {
                        setShowForm(false)
                        onClose && onClose(changed)
                    }}
                    onChange={onChange}
                />
            )}
        </Box>
    )
}

function siteIsPublic(site: ProjectScreenPublicSiteFragment): boolean {
    return site.__typename === "PublicSite" && site.status === PublicSiteStatus.Public
}

function isVisualizationsAllowed(site: ProjectScreenPublicSiteFragment): boolean {
    if (site.__typename === "PublicSite") {
        return site.settings.visualizations
    }

    return true
}

function isAlbumsAllowed(site: ProjectScreenPublicSiteFragment): boolean {
    if (site.__typename === "PublicSite") {
        return site.settings.albums
    }

    return true
}

function EditForm({
    projectId,
    publicSite,
    onChange,
    onClickClose,
    ...props
}: {
    projectId: string
    publicSite: ProjectScreenPublicSiteFragment
    onChange?: () => void
    onClickClose?: (changed: boolean) => void
} & LayerExtendedProps) {
    const [changed, setChanged] = useState(false)

    const [isPublic, setIsPublic] = useState(siteIsPublic(publicSite))

    const [makeProjectPublic, { data: makeProjectPublicResult }] = useMakeProjectPublic(projectId)

    useEffect(() => {
        switch (makeProjectPublicResult?.makeProjectPublic.__typename) {
            case "ProjectMadePublic":
                onChange && onChange()
                setChanged(true)
                break
            case "ProjectIsAlreadyPublic":
                setChanged(!siteIsPublic(publicSite))
                break
            case "Forbidden":
                setIsPublic(siteIsPublic(publicSite))
                setErrorMessage("Доступ запрещен")
                break
        }
    }, [makeProjectPublicResult])

    const [makeProjectNotPublic, { data: makeProjectNotPublicResult }] = useMakeProjectNotPublic(projectId)

    useEffect(() => {
        switch (makeProjectNotPublicResult?.makeProjectNotPublic.__typename) {
            case "ProjectMadeNotPublic":
                onChange && onChange()
                setChanged(true)
                break
            case "ProjectIsAlreadyNotPublic":
                setChanged(!siteIsPublic(publicSite))
                break
            case "Forbidden":
                setIsPublic(siteIsPublic(publicSite))
                setErrorMessage("Доступ запрещен")
                break
        }
    }, [makeProjectNotPublicResult])

    const handleClickPublic = ({ target }: React.ChangeEvent<HTMLInputElement>) => {
        setErrorMessage(undefined)
        setIsPublic(target.checked)

        target.checked ? makeProjectPublic() : makeProjectNotPublic()
    }

    const [errorMessage, setErrorMessage] = useState<string | undefined>()

    return (
        <Layer {...props}>
            <Box pad="medium" gap="medium">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">
                        Ссылка на проект
                    </Heading>
                    <Button
                        icon={<FormClose />}
                        onClick={() => {
                            onClickClose && onClickClose(changed)
                        }}
                    />
                </Box>

                {errorMessage && (
                    <Box
                        pad="small"
                        round="small"
                        direction="row"
                        gap="small"
                        align="center"
                        background={{ color: "status-critical", opacity: "weak" }}
                    >
                        <Box border={{ color: "status-critical", size: "small" }} round="large">
                            <FormClose color="status-critical" size="medium" />
                        </Box>
                        <Text weight="bold" size="medium">
                            {errorMessage}
                        </Text>
                    </Box>
                )}

                <Form>
                    <FormField>
                        <CheckBox
                            label="Сделать проект доступным по ссылке"
                            checked={isPublic}
                            onChange={handleClickPublic}
                        />
                    </FormField>

                    <Accordion>
                        <AccordionPanel label="Настроить доступ">
                            <Box gap="small" pad={{ horizontal: "small", bottom: "small" }}>
                                <Box direction="row" gap="medium">
                                    <CheckBox
                                        label="Визуализации"
                                        checked={isVisualizationsAllowed(publicSite)}
                                        disabled
                                    />
                                </Box>
                                <Box direction="row" gap="medium">
                                    <CheckBox label="Альбом" checked={isAlbumsAllowed(publicSite)} disabled />
                                </Box>
                            </Box>
                        </AccordionPanel>
                    </Accordion>

                    {publicSite.__typename === "PublicSite" && (
                        <FormField label="Ссылка" margin={{ top: "small" }}>
                            <TextInput value={publicSite.url} width="medium" />
                        </FormField>
                    )}
                </Form>

                <Box direction="row" margin={{ top: "medium" }}>
                    <Button type="submit" label="Закрыть" onClick={() => onClickClose && onClickClose(changed)} />
                </Box>
            </Box>
        </Layer>
    )
}
