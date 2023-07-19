// import React from "react";
import { Footer } from '@/Widgets'
import { withProviders } from './providers/index'

import { Routing } from '@/Pages'
import { ContactUs } from '@/Features'

const App = () => {
  return (
    <>
      <Footer />
      <ContactUs />
    </>
  )
}

export default withProviders(App)
