import { Flame } from "lucide-react";
import Categories from "../components/Categories";
import Footer from "../components/Footer";
import GameSlider from "../components/GameSlider";
import Navbar from "../components/Navbar";

export default function () {
    return (
        <div className="flex flex-col bg-primary gap-4">
            <Navbar />

            <main className="flex flex-col text-white grow items-center gap-12">
                <div className="flex flex-col text-center text-2xl font-bold gap-2">
                    <h1>Rastreie os jogos que você já terminou</h1>
                    <h1>Salve os que você quer jogar</h1>
                    <h1>Compartilhe com seus amigos o que é bom</h1>
                </div>

                <button className="bg-green-600 text-white font-bold flex flex-row p-3 rounded-2xl">
                    Comece Agora
                </button>

                <GameSlider title="Destaques" icon={<Flame />} />

                <Categories />

                <GameSlider title="Lançamentos" icon={<Flame />} />

                <GameSlider title="Melhores Avaliados" icon={<Flame />} />
            </main>

            <Footer />
        </div>
    );
}
