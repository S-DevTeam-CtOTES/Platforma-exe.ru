import React from 'react'
import './Footer.scss'
import IconTelegram from '@Icons/IconTelegram.png'
import IconVK from '@Icons/IconVK.png'
import IconYouTube from '@Icons/IconYouTube.png'


const Footer = () => {
  return (
    <footer className='footer'>
      <div className='footer__block'>
        <div className='footer__platforma platforma'>
            <p className='platforma__title'>Platforma.exe</p>
            <div className='platforma__subtitle'>POWERED BY ITC TECHNOSOFT</div>
        </div>
        <div className='footer__contacts contacts'>
            <div className='contacts__title'>Контакты</div>
            <div className='contacts__links'>
                <a className='contacts__link' href="">
                    <img className='contacts__link_icon' src={IconTelegram}/>
                    <p className='contacts__link_title'>Telegram</p>
                </a>
                <a className='contacts__link' href="">
                    <img className='contacts__link_icon' src={IconVK}/>
                    <p className='contacts__link_title'>VK</p>
                </a>
                <a className='contacts__link' href="">
                    <img className='contacts__link_icon' src={IconYouTube}/>
                    <p className='contacts__link_title'>YouTube</p>
                </a>
            </div>
        </div>
      </div>
      <div className='footer__signature'>© Platforma 2023</div>
    </footer>
  )
}

export { Footer }
