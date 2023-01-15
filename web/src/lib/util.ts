export function getValue(obj: object, name: string): any {
    const split = name.split('.')
    let val: Object | string | null = obj
    split.forEach(el => {
        el = el.trim()
        type key = keyof typeof val

        if (val != null && typeof val == 'object' && el in val) {
            val = val[el as key]
        }
    })
    if (val == null) val = ''
    return val
}