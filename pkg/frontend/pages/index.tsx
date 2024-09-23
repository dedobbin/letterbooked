import Head from 'next/head'

export async function getServerSideProps() {
  const res = await fetch("http://localhost:8080/api.json").then(x => x.json());
  return {
    props: {
      status: res.status,
    }
  }
}

export default function Home({status}) {
  return (
    <div>
      <Head>
        <title>letterbooked</title>
        <meta name="description" content="letterbooked" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main >
        <h1>
          books books books
        </h1>
        <div>server status is {status}</div>
      </main>
    </div>
  )
}