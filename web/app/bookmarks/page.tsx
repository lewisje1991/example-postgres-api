'use client'

import { createClient } from '@/utils/supabase/client'
import { useEffect, useState } from 'react'

export default function Page() {
    const [bookmarks, setBookmarks] = useState<any[] | null>(null)
    const supabase = createClient()

    useEffect(() => {
        const getData = async () => {
            const { data } = await supabase.from('bookmarks').select()
            setBookmarks(data)
        }
        getData()
    }, [])

    return <pre>{JSON.stringify(bookmarks, null, 2)}</pre>
}