import { SplitCover } from "features/album/cover/SplitCover/SplitCover"

export function AlbumCover() {
    return (
        <SplitCover
            title="Мылзавод"
            subtitle="Дизайн-проект интерьера квартиры 87&nbsp;м²<br/>в ЖК &laquo;Мылзавод&raquo;"
            qrcodeSrc="http://localhost:8010/qr"
            city="Новосибирск"
            year={2025}
            image={{ __typename: "File", url: "http://localhost:9000/apartomat/p/K0UkDwTV8JUU792OKKojA/i2TtzpZ8F2D6Ax3hvpy43.jpg" }}
            scale={1.0}
            logoText="PUHOVA"
        />
    )
}
