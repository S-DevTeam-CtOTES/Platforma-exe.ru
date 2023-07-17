import { Header, MainInfo, Exploration } from "@/Widgets";
import './Main.scss'


const Main = () => {
  return (
    <div className="Main">
      <div className="Main__triangle"></div>
      <Header />
      <MainInfo/>

      <Exploration/>
      <Exploration/>
      <Exploration/>
    </div>
  );
};

export { Main };
