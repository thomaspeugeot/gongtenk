// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { XLFileDB } from '../xlfile-db'
import { XLFileService } from '../xlfile.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-xlfilestable',
  templateUrl: './xlfiles-table.component.html',
  styleUrls: ['./xlfiles-table.component.css'],
})
export class XLFilesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of XLFile instances
  selection: SelectionModel<XLFileDB> = new (SelectionModel)
  initialSelection = new Array<XLFileDB>()

  // the data source for the table
  xlfiles: XLFileDB[] = []
  matTableDataSource: MatTableDataSource<XLFileDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.xlfiles
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (xlfileDB: XLFileDB, property: string) => {
      switch (property) {
        case 'ID':
          return xlfileDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return xlfileDB.Name;

        case 'NbSheets':
          return xlfileDB.NbSheets;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (xlfileDB: XLFileDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the xlfileDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += xlfileDB.Name.toLowerCase()
      mergedContent += xlfileDB.NbSheets.toString()

      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private xlfileService: XLFileService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of xlfile instances
    public dialogRef: MatDialogRef<XLFilesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.xlfileService.XLFileServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getXLFiles()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "NbSheets",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "NbSheets",
      ]
      this.selection = new SelectionModel<XLFileDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getXLFiles()
    this.matTableDataSource = new MatTableDataSource(this.xlfiles)
  }

  getXLFiles(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.xlfiles = this.frontRepo.XLFiles_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let xlfile of this.xlfiles) {
            let ID = this.dialogData.ID
            let revPointer = xlfile[this.dialogData.ReversePointer as keyof XLFileDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(xlfile)
            }
            this.selection = new SelectionModel<XLFileDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, XLFileDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as XLFileDB[]
          for (let associationInstance of sourceField) {
            let xlfile = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as XLFileDB
            this.initialSelection.push(xlfile)
          }

          this.selection = new SelectionModel<XLFileDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.xlfiles
      }
    )
  }

  // newXLFile initiate a new xlfile
  // create a new XLFile objet
  newXLFile() {
  }

  deleteXLFile(xlfileID: number, xlfile: XLFileDB) {
    // list of xlfiles is truncated of xlfile before the delete
    this.xlfiles = this.xlfiles.filter(h => h !== xlfile);

    this.xlfileService.deleteXLFile(xlfileID).subscribe(
      xlfile => {
        this.xlfileService.XLFileServiceChanged.next("delete")
      }
    );
  }

  editXLFile(xlfileID: number, xlfile: XLFileDB) {

  }

  // display xlfile in router
  displayXLFileInRouter(xlfileID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongxlsx_go-" + "xlfile-display", xlfileID])
  }

  // set editor outlet
  setEditorRouterOutlet(xlfileID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongxlsx_go_editor: ["github_com_fullstack_lang_gongxlsx_go-" + "xlfile-detail", xlfileID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(xlfileID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongxlsx_go_presentation: ["github_com_fullstack_lang_gongxlsx_go-" + "xlfile-presentation", xlfileID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.xlfiles.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.xlfiles.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<XLFileDB>()

      // reset all initial selection of xlfile that belong to xlfile
      for (let xlfile of this.initialSelection) {
        let index = xlfile[this.dialogData.ReversePointer as keyof XLFileDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(xlfile)

      }

      // from selection, set xlfile that belong to xlfile
      for (let xlfile of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = xlfile[this.dialogData.ReversePointer as keyof XLFileDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(xlfile)
      }


      // update all xlfile (only update selection & initial selection)
      for (let xlfile of toUpdate) {
        this.xlfileService.updateXLFile(xlfile)
          .subscribe(xlfile => {
            this.xlfileService.XLFileServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, XLFileDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedXLFile = new Set<number>()
      for (let xlfile of this.initialSelection) {
        if (this.selection.selected.includes(xlfile)) {
          // console.log("xlfile " + xlfile.Name + " is still selected")
        } else {
          console.log("xlfile " + xlfile.Name + " has been unselected")
          unselectedXLFile.add(xlfile.ID)
          console.log("is unselected " + unselectedXLFile.has(xlfile.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let xlfile = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as XLFileDB
      if (unselectedXLFile.has(xlfile.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<XLFileDB>) = new Array<XLFileDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          xlfile => {
            if (!this.initialSelection.includes(xlfile)) {
              // console.log("xlfile " + xlfile.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + xlfile.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = xlfile.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = xlfile.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("xlfile " + xlfile.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<XLFileDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}