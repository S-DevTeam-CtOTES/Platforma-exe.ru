
import './Header.scss'

const Header = () => {
  return (
    <header className='header'>
      <div className="container">
        <nav className="header__nav">
          <div className="header__nav-wrapper">
            <div className="header__nav-wrapper-burger">
              <div className='header__nav-wrapper-burger-line'></div>
              <div className='header__nav-wrapper-burger-line'></div>
            </div>
            <div className="header__nav-wrapper-title">
              Platforma
            </div>
          </div>

          <label className="switch">
              <input type="checkbox" checked />
              <span className="slider round"></span>
          </label>
        </nav> 
      </div>
    </header>
  )
}

export {Header}