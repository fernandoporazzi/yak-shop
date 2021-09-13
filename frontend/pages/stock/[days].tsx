import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../../styles/Home.module.css'
import { GetServerSideProps } from 'next'

const Stock: NextPage = ({ data }) => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Herd</title>
        <meta name="description" content="Get stock info" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Stock data
        </h1>

        <div className={styles.grid}>
          <div className={styles.card}>
            <h2>Milk &rarr;</h2>
            <p>{data.milk}</p>
          </div>
          <div className={styles.card}>
            <h2>Skins &rarr;</h2>
            <p>{data.skins}</p>
          </div>
        </div>
      </main>

      <footer className={styles.footer}>
        <a href="../">
          The YakShop
        </a>
      </footer>
    </div>
  )
}

// This gets called on every request
// export async function getServerSideProps(context) {
export const getServerSideProps: GetServerSideProps = async (context) => {

  // Fetch data from external API
  const res = await fetch(`http://localhost:8080/yak-shop/stock/${context.params?.days}`)
  const data = await res.json()

  // Pass data to the page via props
  return { props: { data } }
}

export default Stock
