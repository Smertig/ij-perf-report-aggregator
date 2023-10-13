import { Observable } from "rxjs"
import { Ref, shallowRef } from "vue"
import { refToObservable } from "../../../configurators/rxjs"
import { PersistentStateManager } from "../PersistentStateManager"

export class MeasureConfiguratorForComparing {
  readonly data = shallowRef<string[] | null>(null)
  private readonly _selected = shallowRef<string[] | string | null>(null)
  readonly state = {
    loading: true,
    disabled: true,
  }

  initData(value: string[]) {
    this.data.value = value
    this._selected.value = value
    this.state.loading = false
    this.state.disabled = false
  }

  createObservable(): Observable<unknown> {
    return refToObservable(this.selected, true)
  }

  setSelected(value: string[] | string | null) {
    this._selected.value = value
  }

  get selected(): Ref<string[] | null> {
    const ref = this._selected
    if (typeof ref.value === "string") {
      ref.value = [ref.value]
    }
    return ref as Ref<string[] | null>
  }

  constructor(measureName: string, persistentStateManager: PersistentStateManager) {
    persistentStateManager.add(measureName, this._selected)
  }
}
