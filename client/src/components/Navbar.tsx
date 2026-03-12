import Logo from "./Logo";

export default function () {
    return (
        <nav
            className="flex flex-row bg-navbar/50 backdrop-blur-xl 
                       justify-center items-center p-4 gap-8 border-b"
        >
            <Logo />

            <ul className="flex flex-row gap-8 font-semibold text-white">
                <li>
                    <p>Entrar</p>
                </li>
                <li>
                    <p>Criar Conta</p>
                </li>
                <li>
                    <p>Jogos</p>
                </li>
                <li>
                    <p>Membros</p>
                </li>
            </ul>
        </nav>
    );
}
