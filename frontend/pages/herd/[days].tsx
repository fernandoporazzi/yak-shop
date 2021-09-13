import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../../styles/Home.module.css'
import { GetServerSideProps } from 'next'

const Herd: NextPage = ({ data }) => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Herd</title>
        <meta name="description" content="Get herd info" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Herd data
        </h1>

        <div className={styles.grid}>
          {data.herd.map((yak, index) => (
            <div key={index} className={styles.card}>
              <h2>{yak.name} &rarr;</h2>
              <p>Age: {yak.age}</p>
              <p>Sex: {yak.sex}</p>
              <p>Last shaved: {yak["age-last-shaved"]}</p>
            </div>
          ))}
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
  const res = await fetch(`http://localhost:8080/yak-shop/herd/${context.params?.days}`)
  const data = await res.json()

  // Pass data to the page via props
  return { props: { data } }
}

export default Herd
