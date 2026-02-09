import React, {useEffect, useRef, useState} from "react"
import {useParams, useNavigate} from "react-router-dom"

import {Box, Button, Grid, Heading, Text, Main, Toolbar, Drop} from "grommet"
import {Print, Close} from "grommet-icons"

import useAlbum, {
  AlbumScreenVisualizationFragment,
  AlbumScreenProjectFragment,
  AlbumScreenAlbumPageCoverFragment,
  AlbumScreenAlbumPageVisualizationFragment,
  AlbumScreenSettingsFragment,
  PageSize as PageSizeEnum,
  PageOrientation as PageOrientationEnum,
  AlbumScreenHouseRoomFragment,
} from "./useAlbum"

import {PageSize, PageOrientation} from "./Settings/"
import AddVisualizations from "pages/Album/AddVisualizations/AddVisualizations"
import {GenerateFile as GenerateAlbumFile} from "pages/Album/GenerateFile"
import {UploadCover} from "pages/Album/UploadCover"
import {AddSplitCover} from "./AddSplitCover/AddSplitCover"
import {AddButtons} from "./AddButtons"
import {Page} from "pages/Album/Page"
import {AlbumScreenAlbumFragment} from "api/graphql"

export function Album() {
  const {id} = useParams<"id">() as { id: string }

  const [visualizations, setVisualizations] = useState<AlbumScreenVisualizationFragment[]>([])

  const [rooms, setRooms] = useState<AlbumScreenHouseRoomFragment[]>([])

  const [settings, setSettings] = useState<AlbumScreenSettingsFragment | undefined>()

  const {
    data,
    loading,
    refetch,
    extracted: {album, project, pages},
  } = useAlbum({id})

  const [showAddVisualizations, setShowAddVisualizations] = useState(false)

  const [showUploadCover, setShowUploadCover] = useState(false)

  const [showAddSplitCover, setShowAddSplitCover] = useState(false)

  const [scale, setScale] = useState(1.0)

  useEffect(() => {
    if (data?.album?.__typename === "Album") {
      if (data?.album?.project?.__typename === "Project") {
        if (data.album.project.visualizations.list.__typename === "ProjectVisualizationsList") {
          setVisualizations(data.album.project.visualizations.list.items)
        }

        if (
          data.album.project.houses.__typename === "ProjectHouses" &&
          data.album.project.houses.list.__typename === "ProjectHousesList"
        ) {
          const list = data.album.project.houses.list.items[0].rooms.list

          if (list.__typename === "HouseRoomsList") {
            setRooms(list.items)
          }
        }
      }

      setSettings(data.album.settings)

      switch (data.album.settings.pageSize) {
        case PageSizeEnum.A3:
          setScale(0.5)
          break
        case PageSizeEnum.A4:
          setScale(0.7)
          break
      }
    }
  }, [data])

  const navigate = useNavigate()

  return (
    <Main overflow="scroll" style={{position: "fixed", inset: 0}} background="background-contrast">
      <Box
        style={{
          position: "fixed",
          top: 10,
          left: 10,
          zIndex: 1,
        }}
        border={{color: "background-front", size: "medium"}}
        round="large"
        background="background-front"
        direction="row"
        align="center"
        pad={{horizontal: "small"}}
        gap="large"
      >
        <Box>
          <Text size="large" weight="bold" onClick={() => project && navigate(`/p/${project.id}`)}>
            {project?.name}
          </Text>
        </Box>
        <Toolbar align="center">
          <Button
            icon={<Print size="medium"/>}
          />
          <Box gap="small" align="end">
            {album && <GenerateAlbumFile album={album} onAlbumFileGenerated={() => refetch()}/>}
          </Box>
        </Toolbar>
      </Box>

      <Box
        style={{
          position: "fixed",
          top: 10,
          right: 10,
          zIndex: 1,
        }}
        border={{color: "transparent", size: "medium"}}
      >
        <Button
          icon={<Close/>}
          onClick={() => {
            project && navigate(`/p/${project.id}`)
          }}
        />
      </Box>

      <Box
        style={{
          position: "fixed",
          top: "84px",
          right: "60px",
          zIndex: 1,
        }}
        background="background-back"
        pad="medium"
        round="xsmall"
        margin={{top: "large"}}
      >
        {settings && (
          <Box gap="small" align="end">
            <Heading level={5} margin={{top: "none"}}>
              Настройки для печати
            </Heading>
            <PageSize albumId={id} size={settings.pageSize} onAlbumPageSizeChanged={() => refetch()}/>
            <PageOrientation
              albumId={id}
              orientation={settings.pageOrientation}
              onAlbumPageOrientationChanged={() => refetch()}
            />
          </Box>
        )}
      </Box>

      <AddButtons
        style={{
          position: "fixed",
          bottom: 0,
          left: 0,
          right: 0,
          zIndex: 1,
        }}
        direction="row"
        justify="center"
        pad="small"
        onClickAddVisualizations={() => setShowAddVisualizations(true)}
        onClickUploadCover={() => setShowUploadCover(true)}
        onClickAddSplitCover={() => setShowAddSplitCover(true)}
      />

      {pages.length > 0 && settings && (
        <Box align="center" pad={{top: "90px", bottom: "68px"}}>
          <Grid width="100%">
            {pages.map((p, pageNumber) => {
              return (
                <Page
                  key={pageNumber}
                  pageNumber={pageNumber}
                  albumId={id}
                  page={p}
                  settings={settings}
                  scale={scale}
                  onPageDeleted={() => {
                    refetch()
                  }}
                />
              )
            })}
          </Grid>
        </Box>
      )}

      {showAddVisualizations && (
        <AddVisualizations
          albumId={id}
          visualizations={visualizations}
          rooms={rooms}
          alreadyAdded={ids(pages)}
          onVisualizationsAdded={() => {
            setShowAddVisualizations(false)
            refetch()
          }}
          onEsc={() => setShowAddVisualizations(false)}
          onClickOutside={() => setShowAddVisualizations(false)}
          onClickClose={() => setShowAddVisualizations(false)}
        />
      )}

      {showUploadCover && (
        <UploadCover
          albumId={id}
          onClickClose={() => setShowUploadCover(false)}
          onAlbumCoverUploaded={async () => {
            setShowUploadCover(false)
            await refetch()
          }}
        />
      )}

      {showAddSplitCover && (
        <AddSplitCover
          albumId={id}
          onClickClose={() => setShowAddSplitCover(false)}
          onSplitCoverAdded={async () => {
            setShowAddSplitCover(false)
            await refetch()
          }}
        />
      )}
    </Main>
  )
}

function orientationWidth(size: PageSizeEnum, orientation: PageOrientationEnum, scale: number = 1.0): string {
  const width = {
    [PageSizeEnum.A3]: {
      [PageOrientationEnum.Portrait]: 297,
      [PageOrientationEnum.Landscape]: 420,
    },
    [PageSizeEnum.A4]: {
      [PageOrientationEnum.Portrait]: 210,
      [PageOrientationEnum.Landscape]: 297,
    },
  }

  return `${width[size][orientation] * scale}mm`
}

function orientationHeight(
  size: PageSizeEnum,
  orientation: PageOrientationEnum = PageOrientationEnum.Landscape,
  scale: number = 1.0
): string {
  const height = {
    [PageSizeEnum.A3]: {
      [PageOrientationEnum.Portrait]: 420,
      [PageOrientationEnum.Landscape]: 297,
    },
    [PageSizeEnum.A4]: {
      [PageOrientationEnum.Portrait]: 297,
      [PageOrientationEnum.Landscape]: 210,
    },
  }

  return `${height[size][orientation] * scale}mm`
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ids(pages: (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]): string[] {
  const vis = pages.filter(
    (p) => p.__typename === "AlbumPageVisualization" && p.visualization.__typename === "Visualization"
  ) as { visualization: { __typename?: "Visualization"; id: string; file: { __typename?: "File"; url: any } } }[]

  return vis.map((v) => v.visualization.id)
}