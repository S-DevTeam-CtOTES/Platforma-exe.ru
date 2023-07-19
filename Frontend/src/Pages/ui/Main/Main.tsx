<<<<<<< HEAD
import { Header, MainInfo, Exploration } from '@/Widgets'
import './Main.scss'
=======
import { Header, MainInfo, Exploration } from "@/Widgets";
import "./Main.scss";
>>>>>>> a4e414beb5ac7af3f9b1fc44941176cc388f3de3

const Main = () => {
  return (
    <div className='Main'>
      <div className='Main__triangle'></div>
      <Header />
      <MainInfo />
      <Exploration />
    </div>
  )
}

export { Main }
